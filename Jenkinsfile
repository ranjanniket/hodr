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
        stage("Docker Build ") {
            steps {
                script {
                    withDockerRegistry(credentialsId: 'docker', toolName: 'docker') {
                        sh "echo 'Buid Docker Image'"
                        sh "docker build -t niket50/hodr:${BUILD_NUMBER} ."
                    }
                }
            }
        }

        stage("Push Artifacts") {
            steps {
                script {
                    withDockerRegistry(credentialsId: 'docker', toolName: 'docker') {
                        sh "echo 'Push to Registry'"
                        sh "docker push niket50/hodr:${BUILD_NUMBER}"
                    }
                }
            }
        }

        stage("TRIVY") {
            steps {
                sh "trivy image -f json -o trivyimage.txt niket50/hodr:${BUILD_NUMBER}"
            }
        }

        stage('Checkout K8S manifest SCM'){
            steps {
                git credentialsId: 'ranjanniket', 
                url: 'https://github.com/ranjanniket/hodr_manifest.git',
                branch: 'main'
            }
        }

        stage('Update Deployment File') {
            steps {
                script {
                    def gitUrl = "https://ranjanniket:${GITHUB_TOKEN}@github.com/ranjanniket/bran_manifest.git"

                    withCredentials([string(credentialsId: 'ranjanniket', variable: 'GITHUB_TOKEN')]) {
                        sh "cat Kubernetes/hodr.yaml"
                        sh "sed -i 's/niket50\\/hodr:.*/niket50\\/hodr:${BUILD_NUMBER}/' Kubernetes/hodr.yaml"
                        sh "cat Kubernetes/hodr.yaml"
                        sh "git add Kubernetes/hodr.yaml"
                        sh "git commit -m 'Update image tag to ${BUILD_NUMBER}'"
                        sh "git remote -v"
                        sh "git push ${gitUrl} HEAD:main"
                    }
                }
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
                attachmentsPattern: 'trivyfs.txt, trivyimage.txt'
            }
        }
}

