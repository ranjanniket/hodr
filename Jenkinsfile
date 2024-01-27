pipeline {
    agent any
    stages {
        stage('Hello from Jenkins') {
            steps {
                echo 'Hello from Jenkins, this is niket'
            }
        }
    }
    post {
        always {
            emailext attachLog: true,
                subject: "'${currentBuild.result}'",
                body: "Project: ${env.JOB_NAME}<br/>" +
                    "Build Number: ${env.BUILD_NUMBER}<br/>" +
                    "URL: ${env.BUILD_URL}<br/>",
                to: 'niketranjan50@gmail.com', 
                attachmentsPattern: 'trivyfs.txt,trivyimage.txt'
        }
    }
}

