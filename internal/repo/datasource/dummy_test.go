package datasource

import (
	"github.com/PereRohit/test-service/internal/config"
	"testing"
)

func Test_dummyDs_HealthCheck(t *testing.T) {
	type fields struct {
		dummySvc *config.DummyInternalSvc
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Success",
			fields: fields{
				dummySvc: &config.DummyInternalSvc{
					Data: "hello world",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dummyDs{
				dummySvc: tt.fields.dummySvc,
			}
			if got := d.HealthCheck(); got != tt.want {
				t.Errorf("HealthCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
