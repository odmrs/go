# Project Structure layout

cmd -> Application specifc code for the executable application

internal -> non-application-specific code potentially reusable code like helpers and the sql databse models

web/ui/html -> user interface assets used by the web application -> html templates
web/ui/static -> the static files like CSS images

## To get another package in another directory:
import (
    packagePlaceHolder "github.com/odmrs/gosnip/internal"
)
func main:
    packagePlaceHolder.Hello

------------------------
package hello // This function need have the first letter in upper case, to import't
func Hello
