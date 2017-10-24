FROM ibmcom/db2express-c:10.5.0.5-3.10.0
MAINTAINER wang-x@neunn.com
ENV DB2INST1_PASSWORD=12345678
ENV LICENSE=accept
RUN mkdir -p /temp/
COPY . /temp/
RUN tar -C /usr/local -xzf /temp/go1.7.5.linux-amd64.tar.gz && \
    chown -R db2inst1:db2inst1 /usr/local/go/
RUN echo 'export GOPATH=/home/db2inst1/go' >> /home/db2inst1/.bash_profile && \
    echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> /home/db2inst1/.bash_profile
RUN mkdir -p /home/db2inst1/go && \
    mkdir -p /home/db2inst1/go/src && \
    cp -r /temp/bitbucket.org /home/db2inst1/go/src/ && \
    chown -R db2inst1:db2inst1 /home/db2inst1/go
# Building and Running setting. （note : Running setting LD_LIBRARY_PATH env already exists.The following is the Building setting.）
RUN echo 'export DB2HOME=$HOME/sqllib' >> /home/db2inst1/.bash_profile && \
    echo 'export CGO_LDFLAGS=-L$DB2HOME/lib' >> /home/db2inst1/.bash_profile && \
    echo 'export CGO_CFLAGS=-I$DB2HOME/include' >> /home/db2inst1/.bash_profile
RUN rm -rf /temp/
