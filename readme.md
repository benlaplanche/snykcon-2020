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
- pack cli 0.13.1 or greataer from buildpacks.io
- kubernetes - e.g. Docker Desktop or minikube would suffice
- snyk cli 1.393 or greater

### Building the container image

You can either use the container image from DockerHub found at `benlaplanche/snykcon-2020:latest` or you can build it yourself as follows

```
$ pack build snykcon-2020:latest --builder "paketobuildpacks/builder:tiny"
```

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

### Part 1 - Kubernetes development & security

Show our basic web app is working

```
$ go run main.go
$ curl localhost:8080
$ go test
```

Here you can see we're getting our hello world message and we have passing tests

We'd like to deploy this to Kubernetes, let reference the example file from [here](https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/)

Here is one i prepared earlier `deployment.yaml`
Matches the reference file, but with the image & deployment names switched. Also with the addition of a load balancer service

Lets check we can deploy and run this to our local Kubernetes running on Docker Desktop

```
$ kubectl apply -f deployment.yaml
$ kubectl get services
$ curl localhost:8080
```

We should get the same hello snykcon response.

We're practicing CI/CD on this project, so we're going to use GitHub actions to run our application tests
But also we're going to run the security scanning step powered by Snyk.

< show the github workflow file >

Lets commit our code

```
$ git push
```

Check-out the github actions workflow for the security issues

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

now lets use the local CLI to validate we've address this issue

```
$ snyk iac test deployment.yaml
```

We can see that this issue has been resolved.

Now lets quickly resolve the rest, given there are a few issues and this is a demo I've made a patch file to speed up the process.

```
$ patch deployment.yaml deployment.patchfile
```

lets run the CLI again

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

### Part 2 - Terraform development & security

We've now got a secure running kubernetes deployment, but we need an S3 bucket for the next iteration of this application as well.
Lets use terraform to provision an S3 bucket for us.

### Part 3 - Break the build...
