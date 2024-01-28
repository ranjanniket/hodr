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
                        sh "docker build -t niket50/hodr:${BUILD_NUMBER} ."
                        sh "docker push niket50/hodr:${BUILD_NUMBER}"
                    }
                }
            }
        }

        stage("TRIVY") {
            steps {
                sh "trivy image -f json -o trivyimage.txt niket50/hodr:latest"
            }
        }

        stage('Update the manifest file') {
            steps {
                script {
                    def newBuildNumber = "${BUILD_NUMBER}"

                    checkout([
                        $class: 'GitSCM',
                        branches: [[name: '*/main']],
                        userRemoteConfigs: [[url: 'https://github.com/ranjanniket/hodr.git']],
                        doGenerateSubmoduleConfigurations: false,
                        extensions: [[$class: 'CleanBeforeCheckout']],
                        submoduleCfg: [],
                    ])

                    sh "sed -i 's/niket50\/hodr:.*/niket50\/hodr:34/' Kubernetes/hodr.yaml"

                    gitAdd = sh(script: "git add Kubernetes/hodr.yaml", returnStatus: true)
                    if (gitAdd == 0) {
                        sh "git commit -m 'Update image tag to ${newBuildNumber}'"
                        sh "git push origin main"
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


