node {
    stage('Cloning repo') {
        checkout scm
    }
    
    stage('Test Docker') {
        sh 'docker -v'
        sh 'docker ps -a'
    }
}
