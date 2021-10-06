package global

const (
	dburi       = "mongodb+srv://pratiktest:standard@cluster0.f653n.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	dbname      = "blog-application"
	performance = 100
)

var (
	jwtSecret = []byte("blogSecret")
)
