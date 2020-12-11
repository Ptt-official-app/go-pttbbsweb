import setuptools

setuptools.setup(
    name="apidoc",
    version='0.0.1',
    install_requires=[
        "pyyaml==5.3.1",
        "Flask==1.1.2",
        "flask-swagger @ git+https://github.com/chhsiao1981/flask-swagger.git@flaskswagger-with-from-file-keyword#egg=flask-swagger",
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "Operating System :: OS Independent",
    ],
)
