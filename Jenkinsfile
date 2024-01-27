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
                git branch: 'main', url: 'https://github.com/ranjanniket/bran.git'
            }
        }

        stage("Sonarqube Analysis") {
            steps {
                withSonarQubeEnv('sonar-server') {
                    sh '''$SCANNER_HOME/bin/sonar-scanner -Dsonar.projectName=bran \
                    -Dsonar.projectKey=bran'''
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
                                sh "docker build -t niket50/bran:v1 ."
                                sh "docker tag niket50/bran:v1 niket50/bran:latest"
                                sh "docker push niket50/bran:latest"
                            }
                        }
                    }
                }

        stage("TRIVY") {
            steps {
                sh "trivy image -f json -o trivyimage.txt niket50/bran:latest"
            }
        }

        stage('Deploy to container') {
            steps {
                sh 'docker run -d -p 8000:8000 -e SECRET_KEY="django-insecure-8j=hrs#^z0t%#1^89isbgqeddf2_zwzh45rz-=h&u%ze)o3e" -e DEBUG="True" -e ALLOWED_HOSTS="*" niket50/bran:latest'
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
