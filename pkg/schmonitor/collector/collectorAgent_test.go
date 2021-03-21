package collector

import (
	"perch/pkg/schmonitor"
	"reflect"
	"testing"
)

func TestAgentBasicCollector(t *testing.T) {
	tests := []struct {
		name    string
		want    AgentBasicInfo
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	AgentBasicCollector()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AgentBasicCollector()
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentBasicCollector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentBasicCollector() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentExector(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AgentExector(); (err != nil) != tt.wantErr {
				t.Errorf("AgentExector() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBasicCollector_CollectorAgent(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			basic := &BasicCollector{}
			if err := basic.CollectorAgent(); (err != nil) != tt.wantErr {
				t.Errorf("CollectorAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBasicCollector_CollectorRegisterWithOpt(t *testing.T) {
	type args struct {
		server schmonitor.ServerTarget
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			basic := &BasicCollector{}
			if err := basic.CollectorRegisterWithOpt(tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("CollectorRegisterWithOpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileCollector_CollectorAgent(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &FileCollector{}
			if err := file.CollectorAgent(); (err != nil) != tt.wantErr {
				t.Errorf("CollectorAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileCollector_CollectorRegisterWithOpt(t *testing.T) {
	type args struct {
		server schmonitor.ServerTarget
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &FileCollector{}
			if err := file.CollectorRegisterWithOpt(tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("CollectorRegisterWithOpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoggerCollector_CollectorAgent(t *testing.T) {
	type fields struct {
		LogLocation string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &LoggerCollector{
				LogLocation: tt.fields.LogLocation,
			}
			if err := log.CollectorAgent(); (err != nil) != tt.wantErr {
				t.Errorf("CollectorAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoggerCollector_CollectorRegisterWithOpt(t *testing.T) {
	type fields struct {
		LogLocation string
	}
	type args struct {
		server schmonitor.ServerTarget
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &LoggerCollector{
				LogLocation: tt.fields.LogLocation,
			}
			if err := log.CollectorRegisterWithOpt(tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("CollectorRegisterWithOpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
