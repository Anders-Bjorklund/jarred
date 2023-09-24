# jarred

A small tool to **simplify uploading jars from your local Maven repository in ~/.m2 to a remote repository** of your choice.  
  
* It can upload everything in your repository or a subset, based on regex inclusion ( all files / paths looking like... ) or regex exclusion ( everything but files / paths looking like... )  
* If you would like to see the number of uploads before you begin, you can **--dry-run**
* Written in Go and can be compiled to a binary for Windows, Linux and / or MacOS
* Needs to have Maven available on the path. 

Example: upload everything from your local repository ( with current directory / pwd being ~/.m2/repository ) :  
**jarred corporate-private-mvn-site-id https://www.url-to-your-mvn-repo/reponame**
  

Example: only upload some artifacts from Apache ( paths beginning with org/apache )  
**jarred --regex-include "^org/apache" corporate-private-mvn-site-id https://www.url-to-your-mvn-repo/reponame**  
  
  
_corporate-private-mvn-site-id_ would already have to be defined in you settings.xml file in ~/.m2, similar to  
```
  <servers>
    <server>
      <id>corporate-private-mvn-site-id</id>
      <username>jimmy.jarred@corporation.com</username>
      <password>abc123</password>
    </server>
  </servers>
```
