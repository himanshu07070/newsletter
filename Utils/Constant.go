package utils

var (
	SmtpHost = "smtp.gmail.com"
	SmtpPort = "587"
	Mime     = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	Body     = `
<body>
   <p>It consist of goroutines and channels.</p>
   <a href={link}>ioscript</a>
</body>
`
)
