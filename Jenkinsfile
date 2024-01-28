pipeline {
    agent any

     environment {
        SCANNER_HOME = tool 'sonar-scanner'
    }

    stages {
        stage('clean workspace') {
            steps {
                cleanWs()
            }
        }

        stage('Checkout from Git') {
            steps {
                git branch: 'main', url: 'https://github.com/ranjanniket/hodr.git'
            }
        }

        stage("Sonarqube Analysis") {
            steps {
                withSonarQubeEnv('sonar-server') {
                    sh '''$SCANNER_HOME/bin/sonar-scanner -Dsonar.projectName=hodr \
                    -Dsonar.projectKey=hodr'''
                }
            }
        }


        stage('OWASP FS SCAN') {
            steps {
                dependencyCheck additionalArguments: '--scan ./ --disableYarnAudit --disableNodeAudit', odcInstallation: 'DP-Check'
                dependencyCheckPublisher pattern: '**/dependency-check-report.xml'
            }
        }
        stage('TRIVY FS SCAN') {
            steps {
                sh "trivy fs . > trivyfs.txt"
            }
        }
        stage("Docker Build & Push") {
                    steps {
                        script {
                            withDockerRegistry(credentialsId: 'docker', toolName: 'docker') {
                                sh "docker build -t niket50/hodr:v1 ."
                                sh "docker tag niket50/hodr:v1 niket50/hodr:latest"
                                sh "docker push niket50/hodr:latest"
                            }
                        }
                    }
                }

        stage("TRIVY") {
            steps {
                sh "trivy image -f json -o trivyimage.txt niket50/hodr:latest"
            }
        }

        stage('Deploy to container') {
            steps {
                sh 'kubectl apply -f Kubernetes/*'
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
