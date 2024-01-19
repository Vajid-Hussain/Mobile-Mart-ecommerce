pipeline {
    agent any

    stages {
        stage ('Diployment'){
            steps {
                script {
                    echo 'echo "Deploying..."'
                    sh 'sudo systemctl restart mobilemart.service'
                    echo 'echo "Deployment complete."'
                }
            }
        }
    }
    
}