pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'minhuy19999/server_golang'
        DOCKER_TAG = '1.0.0'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/MinhUy9999/server-ass.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                sh 'docker image pull minhuy19999/server_golang:1.0.0'
                sh 'docker container stop server_golang || echo "this container does not exist"'
                sh 'docker network create dev || echo "this network exists"'
                sh 'echo y | docker container prune '

                sh 'docker container run -d --rm --name server_golang -p 3000:3000 --network dev minhuy19999/server_golang:1.0.0'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}