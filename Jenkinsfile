def app

pipeline {
    agent any
    stages {
        stage('Cloning Git') {
            steps {
                checkout scm
            }
        }
        stage('Build Image') {
            script {
                app = docker.build("my-image:${env.BUILD_ID}")
            }
 
        }
    }
}
