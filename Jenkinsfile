pipeline {
    agent any
    tools {
        go 'go 1.15'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        registry = "khalid2212/mini_project" 
        registryCredential = 'dockerhubcreds' 
        dockerImage = '' 
    }
    stages {        
        stage('Pull ScientificCalculator') {
            steps {
                git 'https://github.com/falcon2212/ScientificCalculator.git'
            }
        }
        stage('Pre Build') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
                sh 'go get github.com/stretchr/testify'

            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh "chmod 777 ."
                sh "ls -l"
                sh "pwd"
                sh 'go build src/main.go'

            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet ./...'
                    echo 'Running linting'
                    sh 'golint ./...'
                    echo 'Running test'
                    sh 'go test -v ./...'
                }
            }
        }
        stage('Building Docker Image'){
            steps {
                script {
                    dockerImage = docker.build registry
                }
            }
        }
        stage('Deploying Docker Image'){
            steps {
                script {
                    docker.withRegistry('', registryCredential){
                        dockerImage.push()
                    }
                }
            }
        }
        stage('Removing Docker image from jenkins server'){
            steps { 
                sh "docker rmi $registry:latest" 
            }
        }
    }
}