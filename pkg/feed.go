package pkg

// -------テストデータ

var redis = make(map[string][]int)

func init() {
	redis["algorithmfeed:0-earnest"] = []int{1, 2, 3, 4}
	redis["algorithmfeed:0-popular"] = []int{10, 20, 30, 40, 50, 60, 70}
	redis["algorithmfeed:0-others"] = []int{
		100, 200, 300, 400, 500, 600, 700, 800, 900, 1000,
		1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000,
	}
}

// ------------

type CategorizedFeedConfig struct {
	ContentsRate float64
	RedisKey     string

	feed   []int // ここを実際のfeedの型にする
	limit  int
	offset int
}

// GenerateFeed 人気記事と真面目記事を所定の割合・順序で配置してlimit件返す関数
//              offsetを使ってページングされても配置とかは維持される
func GenerateFeed(configs []*CategorizedFeedConfig, offset, limit int) []int {

	othersOffset, othersLimit := 0, 0
	othersFeed := redis["algorithmfeed:0-others"] // 補完用feed(redisに必ずある前提)

	for _, categorizedFeedConfig := range configs {
		// Redisからそれぞれのfeed読み込む
		categorizedFeedConfig.feed = redis[categorizedFeedConfig.RedisKey]
		// 引数の offset と limit を用いて比率に従ってそれぞれの offset と limit 作成
		categorizedFeedConfig.offset, categorizedFeedConfig.limit = generateLimitAndOffsetByGivenRate(offset, limit, categorizedFeedConfig.ContentsRate)
		if sub := categorizedFeedConfig.offset - len(categorizedFeedConfig.feed); sub > 0 {
			othersOffset += sub
		}
	}

	// ここからのconfigスライスの順番がそのまま記事の分類の順番になる
	feed := make([]int, 0, limit)
	for _, categorizedFeedConfig := range configs {
		othersOffset += othersLimit
		othersLimit = 0
		if len(categorizedFeedConfig.feed) < categorizedFeedConfig.offset {
			othersLimit = categorizedFeedConfig.limit
			feed = append(feed, othersFeed[othersOffset:othersOffset+othersLimit]...)
		} else if len(categorizedFeedConfig.feed) <= categorizedFeedConfig.offset+categorizedFeedConfig.limit {
			othersLimit = categorizedFeedConfig.offset + categorizedFeedConfig.limit - len(categorizedFeedConfig.feed)
			feed = append(feed, categorizedFeedConfig.feed[categorizedFeedConfig.offset:]...)
			feed = append(feed, othersFeed[othersOffset:othersOffset+othersLimit]...)
		} else {
			feed = append(feed, categorizedFeedConfig.feed[categorizedFeedConfig.offset:categorizedFeedConfig.offset+categorizedFeedConfig.limit]...)
		}
	}

	return feed
}

func generateLimitAndOffsetByGivenRate(offset, limit int, rate float64) (int, int) {
	return int(float64(offset) * rate), int(float64(limit) * rate)
}
