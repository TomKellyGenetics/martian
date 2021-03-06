//
// Copyright (c) 2014 10X Genomics, Inc. All rights reserved.
//
// Martian runtime tests.
//

package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/martian-lang/martian/martian/syntax"
	"github.com/martian-lang/martian/martian/util"
)

func ExampleBuildCallSource() {
	src, _ := BuildCallSource(
		"STAGE_NAME",
		MakeLazyArgumentMap(map[string]interface{}{
			"input1": []int{1, 2},
			"input2": "foo",
			"input3": json.RawMessage(`{"foo":"bar"}`),
		}),
		nil,
		&syntax.Stage{
			Node: syntax.NewAstNode(15, &syntax.SourceFile{
				FileName: "foo.mro",
				FullPath: "/path/to/foo.mro",
			}),
			Id: "STAGE_NAME",
			InParams: &syntax.InParams{
				List: []*syntax.InParam{
					{
						Tname:    "int",
						ArrayDim: 1,
						Id:       "input1",
					},
					{
						Tname: "string",
						Id:    "input2",
					},
					{
						Tname: "map",
						Id:    "input3",
					},
				},
			},
		}, nil)
	fmt.Println(src)
	// Output:
	// @include "foo.mro"
	//
	// call STAGE_NAME(
	//     input1 = [
	//         1,
	//         2,
	//     ],
	//     input2 = "foo",
	//     input3 = {
	//         "foo": "bar",
	//     },
	// )
}

// Very basic invoke test.
func TestInvoke(t *testing.T) {
	src := `
stage SUM_SQUARES(
    in  float[] values,
    in  int     threads,
    in  bool    local,
    out float   sum,
    src comp    "stages/sum_squares",
)

stage REPORT(
    in  float[] values,
    in  float   sum,
    src exec    "stages/report",
)

pipeline SUM_SQUARE_PIPELINE(
    in  float[] values,
    out float   sum,
)
{
    call SUM_SQUARES(
        values  = self.values,
        threads = 1,
        local   = false,
    )
    call REPORT(
        values = self.values,
        sum    = SUM_SQUARES.sum,
    )

    return (
        sum = SUM_SQUARES.sum,
    )
}

call SUM_SQUARE_PIPELINE(
    values = [1.0, 2.0, 3.0],
)
`
	if d, err := ioutil.TempDir("", "pipestance"); err != nil {
		t.Error(err)
	} else {
		defer os.RemoveAll(d)
		t.Log("Invoking pipestance in ", d)
		pdir := util.RelPath("..")
		if d, err := os.Open(pdir); err != nil {
			t.Skip(err)
		} else {
			// hold open the directory so it doesn't disappear on us.
			defer d.Close()
		}
		t.Log("Runtime directory is ", pdir)
		jobPath := path.Join(pdir, "jobmanagers")
		if _, err := os.Stat(jobPath); os.IsNotExist(err) {
			t.Log("Creating ", jobPath)
			// test harness runs in temp dir.  Need to make a fake config.json.
			if err := os.MkdirAll(jobPath, 0777); err != nil {
				t.Skip(err)
			}
			if d, err := os.Open(jobPath); err != nil {
				t.Skip(err)
			} else {
				defer d.Close()
			}
			defer os.RemoveAll(jobPath)
		} else if err != nil {
			t.Skip(err)
		}
		cfg := path.Join(jobPath, "config.json")
		if _, err := os.Stat(cfg); os.IsNotExist(err) {
			t.Log("Creating ", cfg)
			if ioutil.WriteFile(cfg, []byte(`{
  "settings": {
    "threads_per_job": 1,
    "memGB_per_job": 1,
    "thread_envs": []
  },
  "jobmodes": {}
}`), 0666); err != nil {
				t.Log(err)
			}
			defer os.Remove(cfg)
		} else if err != nil {
			t.Log(err)
		}
		opts := DefaultRuntimeOptions()
		util.SetupSignalHandlers()
		rt := opts.NewRuntime()
		t.Log("Runtime instantiated.")
		if ps, err := rt.InvokePipeline(src,
			path.Join(d, "src.mro"), "test",
			path.Join(d, "test"), nil, "1.0.0",
			make(map[string]string), nil); err != nil {
			t.Error(err)
		} else if ps == nil {
			t.Errorf("nil pipestance")
		} else if _, err := os.Stat(path.Join(d, "test")); err != nil {
			t.Error(err)
		} else {
			ps.Unlock()
		}
	}
}

func TestGetCallableFrom(t *testing.T) {
	callable, err := GetCallableFrom("MY_STAGE",
		path.Join("stages.mro"), []string{"testdata"})
	if err != nil {
		t.Error(err)
	} else if callable.GetId() != "MY_STAGE" {
		t.Errorf("Expected MY_STAGE, got %q", callable.GetId())
	} else if callable.File().FileName != "stages.mro" {
		t.Errorf("Expected stages.mro, got %q", callable.File().FileName)
	}
	callable, err = GetCallableFrom("MY_STAGE",
		path.Join("sub/stages.mro"), []string{"testdata"})
	if err != nil {
		t.Error(err)
	} else if callable.GetId() != "MY_STAGE" {
		t.Errorf("Expected MY_STAGE, got %q", callable.GetId())
	} else if callable.File().FileName != "sub/stages.mro" {
		t.Errorf("Expected stages.mro, got %q", callable.File().FileName)
	}
}

func TestGetCallable(t *testing.T) {
	atd, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}
	mropaths := []string{"testdata", "testdata/", atd, atd + "/"}
	for i := range mropaths {
		callable, err := GetCallable(mropaths[i:], "MY_STAGE", false)
		if err != nil {
			t.Error(err)
		} else if callable.GetId() != "MY_STAGE" {
			t.Errorf("Expected MY_STAGE, got %q", callable.GetId())
		} else if callable.File().FileName != "stages.mro" {
			t.Errorf("Expected stages.mro, got %q (MROPATH: %s)",
				callable.File().FileName,
				strings.Join(mropaths[i:], ":"))
		}
	}
}
