package pkg

import (
	"reflect"
	"testing"
)

func TestGenerateFeed(t *testing.T) {
	type args struct {
		configs []*CategorizedFeedConfig
		limit   int
		offset  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "config指定なし、つまり全て補完用から取ってくる",
			args: args{
				configs: []*CategorizedFeedConfig{},
				offset: 0,
				limit: 4,
			},
			want: []int{100, 200, 300, 400},
		},
		{
			name: "両方足りてて、全部人系記事＋真面目系記事で取れるパターン",
			args: args{
				configs: []*CategorizedFeedConfig{
					{RedisKey: "algorithmfeed:0-earnest", ContentsRate: 0.5},
					{RedisKey: "algorithmfeed:0-popular", ContentsRate: 0.5},
				},
				offset: 0,
				limit:  4,
			},
			want: []int{1, 2, 10, 20},
		},
		{
			name: "真面目系記事が1つ足りなくて補完、人気系記事は全部取れたパターン",
			args: args{
				configs: []*CategorizedFeedConfig{
					{RedisKey: "algorithmfeed:0-earnest", ContentsRate: 0.5},
					{RedisKey: "algorithmfeed:0-popular", ContentsRate: 0.5},
				},
				offset: 4,
				limit:  6,
			},
			want: []int{3, 4, 100, 30, 40, 50},
		},
		{
			name: "真面目系記事はもうないので全部補完、人気系記事は2つだけ取ってあとは補完パターン",
			args: args{
				configs: []*CategorizedFeedConfig{
					{RedisKey: "algorithmfeed:0-earnest", ContentsRate: 0.5},
					{RedisKey: "algorithmfeed:0-popular", ContentsRate: 0.5},
				},
				offset: 10,
				limit:  10,
			},
			want: []int{200, 300, 400, 500, 600, 60, 70, 700, 800, 900},
		},
		{
			name: "両方もうないので、補完のみを使うパターン",
			args: args{
				configs: []*CategorizedFeedConfig{
					{RedisKey: "algorithmfeed:0-earnest", ContentsRate: 0.5},
					{RedisKey: "algorithmfeed:0-popular", ContentsRate: 0.5},
				},
				offset: 20,
				limit:  10,
			},
			want: []int{1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateFeed(tt.args.configs, tt.args.offset, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
