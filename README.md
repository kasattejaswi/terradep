# TERRADEP
This is the module dependency solution for implementing terraform's modules dependency. Using this, users can now manage dependencies both from local file systems and from remote repositories.
<br>
<br>
A simple yaml dependency file is needed to be written in order to manage all the dependencies. These dependencies are version controlled and will be downloaded locally before running terraform.
<br>
<br>
See dependencyConfigs.yaml for sample dependency file.
The project is under development, so dependency config file may change.


# MODULE REPOSITORY
A remote module repository will also be created where these modules can be uploaded by anyone publically and can be used by other developers, so as to promote code reusability.
<br><br>
For an enterprise environment, these modules can also be uploaded in a VCS like Git, SVN or Mercury.

# COMMAND STRUCTURE
## config
This command updates the configuration of terradep.<br>
This includes operations on adding and removing repositories, updating credential information and encrypting/decrypting and locking/unlocking configuration file, so that unauthorized access can be prevented.<br>
It will also generate a new configuration file in case it is not already present inside $USER/.terradep/

```
Usage:
  terradep config [command]

Available Commands:
  init        Initialize an empty configuration in user's home directory
  set-proxy   Set proxy related configuration

Flags:
  -h, --help   help for config
```

## run
This command will start resolving of dependencies in the recursive manner.<br>
It will read the data from the repositories configured in dependencies file and in terradep configuration.<br>

## validate
It will validate the dependencies file and configuration file written.
