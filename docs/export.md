# Using `wsdeploy` for exporting `OpenWhisk` assets

`wskdeploy export` can be used to export `OpenWhisk` assets previously deployed as a *managed project* via `wskdeploy -m manifest.yml`. 
`wskdeploy export` will create a manifest for the managed project assets and separate manifests for each managed project that this managed 
project depends upon, if such dependencies exist and have been described in `manifest.yml` when the managed project has been initially deployed.
The manifest(s) resulting from executing `wskdeploy export` can be later used for deploying at a different `OpenWhisk` instance.

## Use Cases

### Copy `OpenWhisk` Assets

One common scenario, in which the export feature is useful, is populating a newly installed `OpenWhisk` instance with assets from a 
another `OpenWhisk` instance. One might consider a scenario, in which an `OpenWhisk` instance in installed on premises with another `OpenWhisk`
instance residing in the cloud. Consider, for example using one `OpenWhisk` instance on premises and another one in the cloud (e.g., the 
second `OpenWhisk` instance can be [IBM Cloud Functions](https://console.bluemix.net/openwhisk/)). A fairly common scenario is that a developer
will need to deploy assets from the cloud `OpenWhisk` instance on the on premises one and vice-versa. 

### `OpenWhisk` at the Edge

In a variety of IoT and other edge computing scenarios, such as running Virtual Network Functions (VNF) as `OpenWhisk` actions in "edge Data Centers" 
embedded with 5G-MEDIA infrastructure (as pioneered in [5G-MEDIA EU H2020 project](http://www.5gmedia.eu/), there is a need to distribute 
`OpenWhisk` assets developed centrally in the cloud (e.g., [IBM Cloud Functions](https://console.bluemix.net/openwhisk/)) to multiple 
`OpenWhisk` instances running at the edge Data Centers. Again, `wskdeploy export` is handy as a basic tool that allows to automate this. 
management task.

## Basic Usage by Example

Consider a simple manifest file [manifest_lib1.yml](incubator-openwhisk-wskdeploy/tests/src/integration/export/manifest_lib1.yml) for a sample project `lib1`.
The project contains a single `lib1_package` package that comprise three actions (the code of the action in this simple example is the
same for all three, but the action names are different).


### Step 1: deploy `lib1` as a managed project 

```bash
wskdeploy sync -m tests/src/integration/export/manifest_lib1.yml
```

### Step 2: validate `lib1` deployment

```bash
wsk package get lib1_package
```

```

You should see an output similar to the one shown below.

ok: got package lib1_package
{
    "namespace": "your_namespace",
    "name": "lib1_package",
    "version": "0.0.2",
    "publish": false,
    "annotations": [
        {
            "key": "whisk-managed",
            "value": {
                "file": "/root/go_projects/src/github.com/apache/incubator-openwhisk-wskdeploy/tests/src/integration/export/manifest_lib1.yaml",
                "projectDeps": [],
                "projectHash": "80eec5f8e3ee874e22bdacb76aa4cc69aad459c1",
                "projectName": "lib1"
            }
        }
    ],
    "binding": {},
    "actions": [
        {
            "name": "lib1_greeting3",
            "version": "0.0.1",
            "annotations": [
                {
                    "key": "whisk-managed",
                    "value": {
                        "file": "/root/go_projects/src/github.com/apache/incubator-openwhisk-wskdeploy/tests/src/integration/export/manifest_lib1.yaml",
                        "projectDeps": [],
                        "projectHash": "80eec5f8e3ee874e22bdacb76aa4cc69aad459c1",
                        "projectName": "lib1"
                    }
                },
                {
                    "key": "exec",
                    "value": "nodejs:6"
                }
            ]
        },
        {
            "name": "lib1_greeting2",
            "version": "0.0.1",
            "annotations": [
                {
                    "key": "whisk-managed",
                    "value": {
                        "file": "/root/go_projects/src/github.com/apache/incubator-openwhisk-wskdeploy/tests/src/integration/export/manifest_lib1.yaml",
                        "projectDeps": [],
                        "projectHash": "80eec5f8e3ee874e22bdacb76aa4cc69aad459c1",
                        "projectName": "lib1"
                    }
                },
                {
                    "key": "exec",
                    "value": "nodejs:6"
                }
            ]
        },
        {
            "name": "lib1_greeting1",
            "version": "0.0.1",
            "annotations": [
                {
                    "key": "whisk-managed",
                    "value": {
                        "file": "/root/go_projects/src/github.com/apache/incubator-openwhisk-wskdeploy/tests/src/integration/export/manifest_lib1.yaml",
                        "projectDeps": [],
                        "projectHash": "80eec5f8e3ee874e22bdacb76aa4cc69aad459c1",
                        "projectName": "lib1"
                    }
                },
                {
                    "key": "exec",
                    "value": "nodejs:6"
                }
            ]
        }
    ]
}
```

