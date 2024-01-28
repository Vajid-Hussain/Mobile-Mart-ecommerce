pipeline {
    agent any

    stages {

        stage ('build'){
            steps {
                script{
                    sh 'pwd'
                    sh 'which go'
                    sh 'where go'
                    sh 'go run cmd/api/main.go'
                    echo 'building...'
                    sh 'make buildDeployment'
                    echo 'build completed'
                }
            }
        }

        stage ('Diployment'){
            steps {
                script {
                    sh 'pwd'
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