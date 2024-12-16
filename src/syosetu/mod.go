package syosetu

type SyosetuList []struct {
	page  int
	title string
}

type NovelAPIResponse []struct {
	AllCount       int    `json:"allcount,omitempty"`
	Title          string `json:"title,omitempty"`
	Ncode          string `json:"ncode,omitempty"`
	UserID         int    `json:"userid,omitempty"`
	Writer         string `json:"writer,omitempty"`
	Story          string `json:"story,omitempty"`
	BigGenre       int    `json:"biggenre,omitempty"`
	Genre          int    `json:"genre,omitempty"`
	Gensaku        string `json:"gensaku,omitempty"`
	Keyword        string `json:"keyword,omitempty"`
	GeneralFirstUp string `json:"general_firstup,omitempty"`
	GeneralLastUp  string `json:"general_lastup,omitempty"`
	NovelType      int    `json:"novel_type,omitempty"`
	End            int    `json:"end,omitempty"`
	GeneralAllNo   int    `json:"general_all_no,omitempty"`
	Length         int    `json:"length,omitempty"`
	Time           int    `json:"time,omitempty"`
	IsStop         int    `json:"isstop,omitempty"`
	IsR15          int    `json:"isr15,omitempty"`
	IsBL           int    `json:"isbl,omitempty"`
	IsGL           int    `json:"isgl,omitempty"`
	IsZankoku      int    `json:"iszankoku,omitempty"`
	IsTensei       int    `json:"istensei,omitempty"`
	IsTenni        int    `json:"istenni,omitempty"`
	GlobalPoint    int    `json:"global_point,omitempty"`
	DailyPoint     int    `json:"daily_point,omitempty"`
	WeeklyPoint    int    `json:"weekly_point,omitempty"`
	MonthlyPoint   int    `json:"monthly_point,omitempty"`
	QuarterPoint   int    `json:"quarter_point,omitempty"`
	YearlyPoint    int    `json:"yearly_point,omitempty"`
	FavNovelCnt    int    `json:"fav_novel_cnt,omitempty"`
	ImpressionCnt  int    `json:"impression_cnt,omitempty"`
	ReviewCnt      int    `json:"review_cnt,omitempty"`
	AllPoint       int    `json:"all_point,omitempty"`
	AllHyokaCnt    int    `json:"all_hyoka_cnt,omitempty"`
	SasieCnt       int    `json:"sasie_cnt,omitempty"`
	Kaiwaritu      int    `json:"kaiwaritu,omitempty"`
	NovelUpdatedAt string `json:"novelupdated_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

type ExtractType int

const (
	syosetu ExtractType = iota
	syosetuList
)
