pipeline {
    agent any
    tools {
        go 'Go 1.15'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Pre Test') {
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
                sh 'cd src'
                sh 'go build -o ../bin/calculator'
                // sh 'go install'
            }
        }

        // stage('Test') {
        //     steps {
        //         withEnv(["PATH+GO=${GOPATH}/bin"]){
        //             echo 'Running vetting'
        //             sh 'go vet .'
        //             echo 'Running linting'
        //             sh 'golint .'
        //             echo 'Running test'
        //             sh 'cd test && go test -v'
        //         }
        //     }
        // }
        
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "khalidudayagiri2212@gmail.com",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    }  
}