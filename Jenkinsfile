pipeline {
    agent any

    stages {

        stage ('build'){
            steps {
                script{
                    echo 'building...'
                    sh 'go build -o ./cmd/api/tmp/deploy ./cmd/api/main.go'
                    echo 'build completed'
                }
            }
        }

        stage ('Diployment'){
            steps {
                script {
                    echo 'echo "Deploying..."'
                    sh 'sudo systemctl restart mobilemart.service'
                    echo 'echo "Deployment complete."'
                    sh 'sudo systemctl status mobilemart.service'
                    echo 'service file restarted , code deployed...'
                }
            }
        }
    }

}