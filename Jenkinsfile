node {
    stage('Cloning repo') {
        checkout scm
    }
    
    stage('Test Docker') {
        docker.build('go:mv-flix')
    }
}
