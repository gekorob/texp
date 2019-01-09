package out

// The Formatter interface has a simple method to format the entire error stack
// according to the category of each entry
type Formatter interface {
	Format(mq LogQueue) string
}
