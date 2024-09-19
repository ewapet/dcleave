package pathfinder_test

import (
	"github.com/ewapet/dpath/dpath"
	"reflect"
	"testing"
)

type PathFinderGetJsonPathTest struct {
	targetPath     []any
	subject        any
	expectedResult any
}

func (g *PathFinderGetJsonPathTest) Run(t *testing.T) {
	actualResult, _ := dpath.JsonGet(g.subject, g.targetPath)
	if !reflect.DeepEqual(g.expectedResult, actualResult) {
		t.Errorf("got %#v; want %#v", actualResult, g.expectedResult)
	}
}

func Test_GivenASubjectWithOneRootProperty_WhenUsingAPathForAMatchingProperty_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p1"},
		subject: map[string]any{
			"p1": "v1",
		},
		expectedResult: "v1",
	}
	test.Run(t)
}

func Test_GivenASubjectWithTwoRootProperties_WhenUsingAPathForAMatchingProperty_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p2"},
		subject: map[string]any{
			"p1": "v1",
			"p2": "v2",
		},
		expectedResult: "v2",
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyOfDepth2_WhenUsingAPathForTheRootProperty_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": "v1.1",
			},
		},
		expectedResult: map[string]any{
			"p1.1": "v1.1",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyOfDepth2_WhenUsingAPathForThePropertyAtDepth2_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p1", "p1.1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": "v1.1",
			},
		},
		expectedResult: "v1.1",
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyOfDepth3_WhenUsingAPathForThePropertyAtDepth1_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
		},
		expectedResult: map[string]any{
			"p1.1": map[string]any{
				"p1.1.1": "v1.1.1",
			},
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyOfDepth3_WhenUsingAPathForThePropertyAtDepth2_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p1", "p1.1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
		},
		expectedResult: map[string]any{
			"p1.1.1": "v1.1.1",
		},
	}
	test.Run(t)
}

func Test_GivenASubjectWithOnePropertyOfDepth3_WhenUsingAPathForThePropertyAtDepth3_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p1", "p1.1", "p1.1.1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
		},
		expectedResult: "v1.1.1",
	}
	test.Run(t)
}

func Test_GivenASubjectHavingPropertiesOfMixedDepth_WhenUsingAPathForARootProperty_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p3"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
			"p2": map[string]any{
				"p2.1": 5,
			},
			"p3": 5,
		},
		expectedResult: 5,
	}
	test.Run(t)
}

func Test_GivenASubjectHavingPropertiesOfMixedDepth_WhenUsingAPathForANestedProperty_ThenTheValueIsReturned(t *testing.T) {
	test := PathFinderGetJsonPathTest{
		targetPath: []any{"p2", "p2.1"},
		subject: map[string]any{
			"p1": map[string]any{
				"p1.1": map[string]any{
					"p1.1.1": "v1.1.1",
				},
			},
			"p2": map[string]any{
				"p2.1": 5,
			},
			"p3": 5,
		},
		expectedResult: 5,
	}
	test.Run(t)
}
