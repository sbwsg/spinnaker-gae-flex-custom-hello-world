pipeline {
  agent any
  stages {
    stage('Fetch') {
      steps {
        git(url: 'https://github.com/sbwsg/spinnaker-gae-flex-custom-hello-world', branch: 'master', poll: true)
      }
    }
    stage('Build') {
      steps {
        sh 'go build -o hello .'
      }
    }
    stage('Deploy') {
      steps {
        fileExists 'hello'
      }
    }
  }
}