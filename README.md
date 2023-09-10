# jarred

A small tool to **simplify uploading jars from your local Maven repository in ~/.m2 to a remote repository** of your choice.  
  
* It can upload everything in your repository or a subset, based on regex inclusion ( all files / paths looking like... ) or regex exclusion ( everything but files / paths looking like... )  
* If you would like to see the nuber of uploads before you begin, you can **--dry-run**
* Written in Go and can be compiled to a binary for Windows, Linux and / or MacOS
* Needs to have Maven available on the path. 

Example: upload everything from local repository ( with current directory / pwd being ~\.m2\repository ) :  
**jarred --site-id corporate-private-mvn**

_corporate-private-mvn_ would already have to be defined in you settings.xml file in ~/.m2, similar to  
```
  <servers>
    <server>
      <id>corporate-private-mvn</id>
      <username>jimmy.jarred@corporation.com</username>
      <password>abc123</password>
    </server>
  </servers>
```
