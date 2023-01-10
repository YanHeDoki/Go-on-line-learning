package dao

type Teacher struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Intro  string `json:"intro" db:"intro"`
	Career string `json:"career" db:"career"`
	Level  int    `json:"level" db:"level"`
	Sort   int    `json:"sort" db:"sort"`
	Avatar string `json:"avatar" db:"avatar"`
}

type Course struct {
	Id          string ` json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	ParentId    string `json:"parent_id" db:"parent_id"`
	Sort        int    `json:"sort" db:"sort"`
	GmtCreate   string `json:"gmt_create" db:"gmt_create"`
	GmtModified string `json:"gmt_modified" db:"gmt_modified"`
}

type CourseTree struct {
	Id       string        ` json:"id" db:"id"`
	Title    string        `json:"title" db:"title"`
	Children []*CourseNode `json:"children"`
}

type CourseNode struct {
	Id       string ` json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	ParentId string `json:"parent_id" db:"parent_id"`
}

type CourseInfo struct {
	Id              string  `json:"id" db:"id"`
	TeacherId       string  `json:"teacherId" db:"teacher_id"`
	SubjectId       string  `json:"subjectId" db:"subject_id"`
	SubjectParentId string  `json:"subjectParentId" db:"subject_parent_id"`
	Title           string  `json:"title" db:"title"`
	Cover           string  `json:"cover" db:"cover"`
	Description     string  `json:"description" db:"description"`
	LessonNum       int     `json:"lessonNum" db:"lesson_num"`
	Price           float64 `json:"price" db:"price"`
	BuyCount        int     `json:"buy_count" db:"buy_count"`
	GmtCreate       string  `json:"gmt_create" db:"gmt_create"`
}

type PublicCourseInfo struct {
	Id              string  `json:"id" db:"id"`
	TeacherName     string  `json:"teacherName" db:"Teacher_name"`
	SubjectLevelOne string  `json:"subjectLevelOne" db:"subject_id"`
	SubjectLevelTwo string  `json:"subjectLevelTwo" db:"subject_parent_id"`
	Title           string  `json:"title" db:"title"`
	Cover           string  `json:"cover" db:"cover"`
	LessonNum       int     `json:"lessonNum" db:"lesson_num"`
	Price           float64 `json:"price" db:"price"`
}

type CourseFront struct {
	Id              string  `json:"id" db:"id"`
	Title           string  `json:"title" db:"title"`
	Cover           string  `json:"cover" db:"cover"`
	Description     string  `json:"description" db:"description"`
	LessonNum       int     `json:"lessonNum" db:"lesson_num"`
	Price           float64 `json:"price" db:"price"`
	BuyCount        int     `json:"buyCount" db:"buy_count"`
	ViewCount       int     `json:"viewCount" db:"view_count"`
	GmtCreate       string  `json:"gmtCreate" db:"gmt_create"`
	TeacherAvatar   string  `json:"teacherAvatar" db:"teacher_avatar"`
	TeacherId       string  `json:"teacherId" db:"teacher_id"`
	TeacherName     string  `json:"teacherName" db:"Teacher_name"`
	SubjectLevelOne string  `json:"subjectLevelOne" db:"subjectLevelOne"`
	SubjectLevelTwo string  `json:"subjectLevelTwo" db:"subjectLevelTwo"`
	SubjectId       string  `json:"subjectId" db:"subject_id"`
	SubjectParentId string  `json:"subjectParentId" db:"subject_parent_id"`
}

type CourseDescription struct {
	Id          string `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}

type EduCourse struct {
	Id              string  `json:"id" db:"id"`
	TeacherId       string  `json:"teacher_id" db:"teacher_id"`
	SubjectId       string  `json:"subject_id" db:"subject_id"`
	SubjectParentId string  `json:"subject_parent_id" db:"subject_parent_id"`
	Title           string  `json:"title" db:"title"`
	Price           float64 `json:"price" db:"price"`
	LessonNum       int     `json:"lesson_num" db:"lesson_num"`
	Cover           string  `json:"cover" db:"cover"`
	BuyCount        int     `json:"buy_count" db:"buy_count"`
	ViewCount       int     `json:"view_count" db:"view_count"`
	Version         int     `json:"version" db:"version"`
	Status          string  `json:"status" db:"status"`
}
type Chapter struct {
	Id       string ` json:"chapterId" db:"id"`
	Title    string `json:"title" db:"title"`
	CourseId string `json:"courseId" db:"course_id"`
	Sort     int    `json:"sort" db:"sort"`
}

type ChapterTree struct {
	Id       string         ` json:"id" db:"id"`
	Title    string         `json:"title" db:"title"`
	Children []*ChapterNode `json:"children"`
}

type ChapterNode struct {
	Id                string ` json:"id" db:"id"`
	Title             string `json:"title" db:"title"`
	ChapterId         string `json:"chapter_id" db:"chapter_id"`
	VideoSourceId     string `json:"videoSourceId" db:"video_source_id"`
	VideoOriginalName string `json:"video_original_name" db:"video_original_name"`
}

type VideoInfo struct {
	Id                string  `json:"id" db:"id"`
	CourseID          string  `json:"courseId" db:"course_id"`
	ChapterId         string  `json:"chapterId" db:"chapter_id"`
	Title             string  `json:"title" db:"title"`
	VideoSourceId     string  `json:"videoSourceId" db:"video_source_id"`
	VideoOriginalName string  `json:"videoOriginalName" db:"video_original_name"`
	Sort              int     `json:"sort" db:"sort"`
	Duration          float32 `json:"duration" db:"duration"`
}
