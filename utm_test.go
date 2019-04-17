package utm

import "testing"

func TestGenerate(t *testing.T) {
	type args struct {
		base     string
		source   string
		medium   string
		campaign string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: struct {
				base     string
				source   string
				medium   string
				campaign string
			}{base: "http://abc.com/", source: "source_test", medium: "medium_test", campaign: "campaign_test"},
			want: "http://abc.com/?utm_campaign=campaign_test&utm_medium=medium_test&utm_source=source_test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.base, tt.args.source, tt.args.medium, tt.args.campaign)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
