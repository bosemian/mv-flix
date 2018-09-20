pipeline {
    agent any
    tools {
        nodejs 'nodejs'
    }
    stages {
        stage('Cloning Git') {
            steps {
                checkout scm
            }
        }
        stage('Initialize') {
            steps {
                sh 'yarn -version'
                sh 'docker -v'
            }
        }
    }
}
