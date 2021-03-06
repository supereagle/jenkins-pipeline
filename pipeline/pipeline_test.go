package pipeline_test

import (
	"testing"

	"github.com/supereagle/goline/api"
	"github.com/supereagle/goline/pipeline"
)

func TestValidatePipeline(t *testing.T) {
	type PipelineValidator struct {
		pipeline *api.Pipeline
		result   bool
	}

	pvs := []PipelineValidator{
		PipelineValidator{
			pipeline: &api.Pipeline{
				Name: "validate-jdk1",
				Jdk:  "1.7",
				Repo: &api.Repo{
					RepoPath: "git@test.com:test/test.git",
					Branch:   "master",
				},
			},
			result: false,
		},
		PipelineValidator{
			pipeline: &api.Pipeline{
				Name: "validate-jdk2",
				Jdk:  "jdk1.8",
				Repo: &api.Repo{
					RepoPath: "git@test.com:test/test.git",
					Branch:   "master",
				},
				ProjectType: "maven",
				Project: api.MavenProject{
					RootPom: "pom.xml",
				},
			},
			result: true,
		},
		PipelineValidator{
			pipeline: &api.Pipeline{
				Name: "validate-period-trigger",
				Jdk:  "jdk1.8",
				Repo: &api.Repo{
					RepoPath: "git@test.com:test/test.git",
					Branch:   "master",
				},
				PeriodTrigger: &api.PeriodTrigger{
					Skipped:  false,
					Strategy: "H/30 * * * *",
				},
				ProjectType: "maven",
				Project: api.MavenProject{
					RootPom: "pom.xml",
				},
			},
			result: true,
		},
		PipelineValidator{
			pipeline: &api.Pipeline{
				Name: "validate-period-trigger2",
				Jdk:  "jdk1.8",
				Repo: &api.Repo{
					RepoPath: "git@test.com:test/test.git",
					Branch:   "master",
				},
				PeriodTrigger: &api.PeriodTrigger{
					Skipped: false,
				},
				ProjectType: "maven",
				Project: api.MavenProject{
					RootPom: "pom.xml",
				},
			},
			result: false,
		},
	}

	for _, pv := range pvs {
		if pipeline.ValidatePipeline(pv.pipeline) != pv.result {
			t.Errorf("Pipeline %s's config is not correct!", pv.pipeline.Name)
		}
	}
}
