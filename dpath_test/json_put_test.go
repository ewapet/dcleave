package pathfinder_test

import (
	"github.com/ewapet/dpath/dpath"
	"reflect"
	"testing"
)

type PutJsonPathTest struct {
	targetPath     []any
	subject        any
	value          any
	expectedResult any
}

func (g *PutJsonPathTest) Run(t *testing.T) {
	dpath.JsonPut(g.subject, g.value, g.targetPath)
	if !reflect.DeepEqual(g.expectedResult, g.subject) {
		t.Errorf("got %#v; want %#v", g.subject, g.expectedResult)
	}
}

func Test_GivenASubjectWithNoProperties_WhenSettingAPropertyAtDepth0_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1"},
		subject:    map[string]any{},
		value:      "v1",
		expectedResult: map[string]any{
			"p1": "v1",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithNoProperties_WhenSettingAPropertyAtDepth1_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1"},
		subject:    map[string]any{},
		value:      "v1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": "v1.1",
			},
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithNoProperties_WhenSettingAPropertyAtDepth2_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1", "p1.1.1"},
		subject:    map[string]any{},
		value:      "v1.1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyAtDepth0_WhenSettingANewPropertyAtDepth0_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1"},
		subject:    map[string]any{"p2": "v2"},
		value:      "v1",
		expectedResult: map[string]any{
			"p1": "v1",
			"p2": "v2",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyAtDepth0_WhenSettingANewPropertyAtDepth1_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1"},
		subject:    map[string]any{"p2": "v2"},
		value:      "v1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": "v1.1",
			},
			"p2": "v2",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyAtDepth0_WhenSettingANewPropertyAtDepth2_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1", "p1.1.1"},
		subject:    map[string]any{"p2": "v2"},
		value:      "v1.1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
			"p2": "v2",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOneExistingPropertyAtDepth0_WhenOverridingAPropertyAtDepth0_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1"},
		subject: map[string]any{
			"p1": "x1",
		},
		value: "v1",
		expectedResult: map[string]any{
			"p1": "v1",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOneExistingPropertyAtDepth0_WhenOverridingAPropertyAtDepth1_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1"},
		subject: map[string]any{
			"p1": "x1",
		},
		value: "v1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": "v1.1",
			},
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOneExistingPropertyAtDepth0_WhenOverridingAPropertyAtDepth2_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1", "p1.1.1"},
		subject: map[string]any{
			"p1": "x1",
		},
		value: "v1.1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithVariableDepth_WhenOverridingProperties_ThenThePropertyIsSet(t *testing.T) {

	test := PutJsonPathTest{
		targetPath: []any{"p1", "p1.1", "p1.1.1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{},
				"p1.2": 42,
			},
		},
		value: "v1.1.1",
		expectedResult: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
				"p1.2": 42,
			},
		},
	}
	test.Run(t)
}
