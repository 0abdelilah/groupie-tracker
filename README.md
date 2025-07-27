RUN:
- go run . 

API:
- /artists -> show all artists
- /locations/{id} -> show location of a given artist
- /relation/{id} -> show relations of a given artist
- /dates/{id} -> show dates of a given artist


todo:
use local json instead of fetching
handle 404