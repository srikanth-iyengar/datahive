FROM apache/spark
USER root

# Install SSH and OpenSSH server
RUN apt-get update && apt-get install -y openssh-server

# Generate SSH key pair for the root user (or a user of your choice)
RUN ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa
RUN cat /root/.ssh/id_rsa.pub >> /root/.ssh/authorized_keys
RUN chmod 600 /root/.ssh/authorized_keys
RUN chown root:root /root/.ssh/authorized_keys

# Configure SSH to allow root login and disable strict host key checking
RUN echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
RUN echo "StrictHostKeyChecking no" >> /etc/ssh/ssh_config

# Start SSH server
RUN service ssh start

# Start Spark master, Spark worker, and Thrift server
RUN printf "#!/bin/bash \n\
service ssh start \n\
/opt/spark/sbin/start-all.sh \n\
tail -f /opt/spark/logs/*" >> /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT [ "/docker-entrypoint.sh" ]

