pipeline {
    agent any
    def app = ''
    stages {
        stage('Cloning Git') {
            steps {
                checkout scm
            }
        }
        stage('Build Image') {
            app = docker.build("test:go")
        }
    }
}
