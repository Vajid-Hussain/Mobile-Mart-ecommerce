pipeline {
    agent any

    environment {
        PATH = '/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin'
    }

    stages {

        stage ('testing'){
            steps{
                script{
                    echo 'start testing....'
                    sh 'make test'
                    echo 'test complite succesfully.. '
                }
            }
        }

        stage ('build'){
            steps {
                script{
                    sh 'pwd'
                    sh 'which go'
                    echo 'take build...'
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