pipeline {
    agent { 
        node {label 'slave-2'} 
    }

    stages {
        stage('Prepare job definition') {
            steps {
                script {
                    def branch = env.BRANCH_NAME
                    def image_name = jenskin + branch
                    echo 'Pulling...' + branch
                    withDockerRegistry(credentialsId: 'd2330c30-d21b-4670-84ed-37dd988b506a') {
                        sh 'docker build --no-cache -t unknopu/${image_name} .'
                        sh 'docker push unknopu/${image_name}'
                    }
                }
            }
            post {
                always { 
                    cleanWs()
                }
            }
        }
    }
}


