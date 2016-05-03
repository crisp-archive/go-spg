package var

type Var struct {
    Name string
    Description string
    Keywords []string
    PageList []Page
    BlogList []Blog
    BlogPerPage uint32
}

type Page struct {
    Title string
    Tags []string
    Timestamp uint32
    FileName string
}

type Blog struct {
    Title string
    Abstract string
    Tags []string
    Timestamp uint32
    FileName string
    Disqus bool
}

func New(name string, desc string, keywords []string, page_list []Page, blog_list []Blog, blog_per_page uint32) {
}

func (var *Var) AddBlog(Blog) {
}

func (var *Var) GetLastBlog() (Blog) {
}

func (var *Var) AddPage(Page) {
}
