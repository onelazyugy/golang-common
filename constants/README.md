# Constants 

Reads common environment variables that are setup and contains any other common constants for the Shasta Project that 
can be used by different teams.

If package A imports Constants, Constants.init() is called before A.init() is called. 

```go  
import "https://github.com/ncr-swt-retail/shasta-common-golib/constants"
```