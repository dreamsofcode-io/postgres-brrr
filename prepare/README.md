# Prepare

This directory can be used to test the performance of prepared statements.

## Setup

- Python 3
- pip
- sqlx-cli


To do so, first load up a virtual environment for python:

```
$ python -m venv .venv
$ source .venv/bin/activate
```

then install the dependencies via pip

```
$ pip install -r requirements.txt
```

With python installed, you can set up the database using the `sqlx-cli`

```
$ sqlx migrate run
```

Finally, you can run the bench.py

```
$ python bench.py
```
