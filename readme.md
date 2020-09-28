# Snykcon 2020 Demo Repo

**Talk Title:** How to deploy securely using Kubernetes & Terraform
**Talk Abstract:**
Kubernetes is fast becoming the platform of choice for deploying modern cloud native applications and Terraform is increasingly the tool of choice for creating infrastructure to support these applications.

Their flexibility means they are powerful for a wide range of use-cases and their focus on configuration in code means they are accessible to development teams to use quickly and autonomously.

But with this comes the challenge of knowing whether you’ve deployed your application securely. How do you understand all of the potential configuration options and their impact? How do you know that the supporting infrastructure is appropriately locked down and you are following your own teams best practices?

In this talk we will:

- Look at a typical development flow for writing and validating a Kubernetes and Terraform deployment, starting from the command line through to your source control system
- Discuss the the challenges and security considerations you should be aware of and how to work with your security team if you have one
- Show a few demos of tools, including Snyk, that can help you get faster feedback

**Audience Takeaways:**
All attendees should come away from the session with some practical ideas they can put into practice straight away, whether they have wide adoption of Kubernetes and Terraform yet or not.

For development teams building applications we’ll look at:

- Why considering security from the beginning is beneficial
- How to securely deploy to Kubernetes and the considerations in doing so
- How to securely provision infrastructure using Terraform
- How to seamless add security into your local development workflow, with the toolchain you are familiar with
- How to work collaboratively with the security team

And for security teams responsible for assuring the applications and infrastructure that is being deployed, we’ll discuss:

- How to get visibility into each application and across applications
- How to engage with development teams to educate and empower them to develop securely

# Demo Flow

## Pre-requisites

- golang 1.14.5
- pack cli 0.13.1 or greater from buildpacks.io
- kubernetes - Docker Desktop or minikube would suffice
- snyk cli 1.402 or greater

### Building the container image

You can either use the container image from DockerHub found at `benlaplanche/snykcon-2020:latest` or you can build it yourself as follows

```
$ pack build snykcon-2020:latest --builder "paketobuildpacks/builder:tiny"
```

This is using [Cloud Native Buildpacks](https://buildpacks.io/)

If you have docker running locally you can quickly testing the image is working by

```
$ docker run --rm -p 8080:8080 snykcon-2020
$ curl localhost:8080
```

If you get a message back saying `Hello Snykcon 2020!!!` you're up and running.

You can then tag and push the image to DockerHub as follows

```
$ docker tag snykcon-2020:latest benlaplanche/snykcon-2020:latest
$ docker login
$ docker push benlaplanche/snykcon-2020:latest
```

## Demo flow

### Part 1 - Deploy to Kubernetes

Show our basic web app is working

```
$ go run main.go
$ curl localhost:8080
```

Terminal shows `Hello Snykcon 2020!!!`

Next, deploy to Kubernetes.
Reference the example file from [here](https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/)

`deployment.yaml` is based on this with the image & deployment names updated.
It also has the addition of a load balancer service.

Deploy to Kubernetes

```
$ kubectl apply -f deployment.yaml
$ kubectl get services
$ curl localhost:8080
```

Terminal shows `Hello Snykcon 2020!!!`

We're using GitHub actions for our CI/CD in this demo.
Show the Github workflow file `./github/workflows/snyk.yaml`

Lets commit our code

```
$ git push
```

Review the GitHub actions workflow output (here)[https://github.com/benlaplanche/snykcon-2020/actions]

Lets address these issues locally

```
$ code deployment.yaml
```

Address the `Container can run as the root [Medium Severity] [SNYK-CC-K8S-10]` issue first
add the following yaml to the containers spec

```
securityContext:
  runAsNonRoot: true
```

Validate we've address this using the CLI

```
$ snyk iac test deployment.yaml
```

We can see that this issue has been resolved.

Quickly address the rest using the patch file

```
$ patch deployment.yaml deployment.patch
```

Run the CLI again

```
$ snyk iac test deployment.yaml
```

We can see that all of the issues have been addressed.

Lets commit these changes and move on

```
$ git commit -am 'Resolved kubernetes issues reported by Snyk'
$ git push
```

You can validate this through the newly triggered GitHub actions as well.

### Part 2 - Provision using Terraform

Deploy an S3 bucket to store our conference pictures in using terraform.

Show `main.tf`

Navigate to (Snyk and import our project)[https://app.snyk.io] from GitHub.
Show the identified issues in the UI.

### Part 3 - Collaborate with Security

Show changing the severity of an issue in the (Infrastructure as Code settings)[https://app.snyk.io/org/snykcon-2020/manage/cloud-config]

Add Terraform scanning to our GitHub actions.

```
$ cd .github/workflows
$ ls
$ patch snyk.yaml snyk.patch
```

Show `.github/workflows/snyk.yaml`

Two key differences
`continue-on-error` is set to `false` which will cause this pipeline step to fail if an issue is identified by Snyk
`--severity-threshold=high` is passed as an additional argument. This filters the results to only show High severity or higher issues.

Lets push our code

```
$ git commit -am 'Added terraform resource for s3 and security checks'
$ gith push
```

Review the GitHub actions workflow output (here)[https://github.com/benlaplanche/snykcon-2020/actions]
Notice that the Terraform pipeline step has failed.

Congratulations, you have reached the end of the demo!

### Appendix

If you'd like to address the Terraform issue to have a passing build
Change the `acl` to be `private` on the `main.tf` file

Commit changes and push

```
$ git commit -am 'restricted acl to be private'
$ git push
```

Check the build is now passing

Congrats! You have a secured infrastructure.
