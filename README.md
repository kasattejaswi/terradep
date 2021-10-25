# TERRADEP

This is the module dependency solution for implementing terraform's modules dependency. Using this, users can now manage dependencies both from local file systems and from remote repositories.
<br>
A simple yaml dependency file is needed to be written in order to manage all the dependencies. These dependencies are version controlled and will be downloaded locally before running terraform.
<br>
See dependencyConfigs.yaml for sample dependency file.
The project is under development, so dependency config file may change.


# MODULE REPOSITORY
A remote module repository will also be created where these modules can be uploaded by anyone publically and can be used by other developers, so as to promote code reusability.
<br>
For an enterprise environment, these modules can also be uploaded in a VCS like Git, SVN or Mercury.