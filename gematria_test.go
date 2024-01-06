package go_gematria

import (
	`reflect`
	`testing`
)

func TestNewGemScore(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want Gematria
	}{
		{
			name: "test the string andrei",
			args: args{
				data: "andrei",
			},
			want: Gematria{
				Jewish:  139,
				English: 306,
				Simple:  51,
			},
		},
		{
			name: "manifesting three six nine",
			args: args{
				data: "manifesting three six nine",
			},
			want: Gematria{
				Jewish:  1028,
				English: 1602,
				Simple:  267,
			},
		},
		{
			name: "Jesus Christ Prince of Peace",
			args: args{
				data: "Jesus Christ Prince of Peace",
			},
			want: Gematria{
				Jewish:  1602,
				English: 1602,
				Simple:  267,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGematria(tt.args.data)
			if err != nil {
				t.Errorf("NewGematria() received error = %v, want %v", err, tt.want)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGematria() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewGemScore(b *testing.B) {
	benchmarks := []struct {
		name string
		data string
	}{
		{
			name: "Benchmark for the string andrei",
			data: "andrei",
		},
		{
			name: "Benchmark for the string manifesting three six nine",
			data: "manifesting three six nine",
		},
		{
			name: "Benchmark for the string Jesus Christ Prince of Peace",
			data: "Jesus Christ Prince of Peace",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = NewGematria(bm.data)
			}
		})
	}
}
