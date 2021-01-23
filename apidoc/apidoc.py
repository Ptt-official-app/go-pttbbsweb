# -*- coding: utf-8 -*-

import yaml
from flask import Flask

app = Flask(__name__.split('.')[0])

_APP_PREFIX = '/api'


def _with_app_prefix(path):
    """Summary

    Args:
        path (TYPE): Description

    Returns:
        TYPE: Description
    """
    if not _APP_PREFIX:
        return path
    return _APP_PREFIX + path


with open('apidoc/template.json', 'r') as f:
    template = yaml.load(f)


@app.route('/')
def _index():
    """
    swagger_from_file: apidoc/index.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/login'), methods=['POST'])
def _login():
    """
    swagger_from_file: apidoc/login.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/register'), methods=['POST'])
def _register_user():
    """
    swagger_from_file: apidoc/register_user.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/attemptregister'), methods=['POST'])
def _attempt_register_user():
    """
    swagger_from_file: apidoc/attempt_register_user.yaml

    """
    return ''


@app.route(_with_app_prefix('/articles/popular'))
def _load_popular_articles():
    """
    swagger_from_file: apidoc/load_popular_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/favorites'))
def _load_favorite_boards(user_id):
    """
    swagger_from_file: apidoc/load_favorite_boards.yaml

    """
    return ''


@app.route(_with_app_prefix('/boards'))
def _load_general_boards():
    """
    swagger_from_file: apidoc/load_general_boards.yaml
    """
    return ''


@app.route(_with_app_prefix('/boards/byclass'))
def _load_general_boards_by_class():
    """
    swagger_from_file: apidoc/load_general_boards_by_class.yaml
    """
    return ''


@app.route(_with_app_prefix('/boards/popular'))
def _load_popular_boards():
    """
    swagger_from_file: apidoc/load_popular_boards.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>'))
def _get_user_info(user_id):
    """
    swagger_from_file: apidoc/get_user_info.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/summary'))
def _get_user_summary(user_id):
    """
    swagger_from_file: apidoc/get_user_summary.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/updatepasswd'), methods=['POST'])
def _change_passwd(user_id):
    """
    swagger_from_file: apidoc/change_passwd.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/articles'))
def _load_user_articles(user_id):
    """
    swagger_from_file: apidoc/load_user_articles.yaml

    Args:
        username (TYPE): Description

    Returns:
        TYPE: Description
    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/comments'))
def _load_user_comments(user_id):
    """
    swagger_from_file: apidoc/load_user_comments.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/attemptchangeemail'), methods=['POST'])
def _attempt_change_email(user_id):
    """
    swagger_from_file: apidoc/attempt_change_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/changeemail'), methods=['POST'])
def _change_email(user_id):
    """
    swagger_from_file: apidoc/change_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/attemptsetidemail'), methods=['POST'])
def _attempt_set_id_email(user_id):
    """
    swagger_from_file: apidoc/attempt_set_id_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<user_id>/setidemail'), methods=['POST'])
def _set_id_email(user_id):
    """
    swagger_from_file: apidoc/set_id_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/articles'))
def _load_general_articles(bid):
    """
    swagger_from_file: apidoc/load_general_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/articles/bottom'))
def _load_bottom_articles(bid):
    """
    swagger_from_file: apidoc/load_bottom_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>'))
def _get_article_detail(bid, aid):
    """
    swagger_from_file: apidoc/get_article_detail.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/comments'))
def _load_article_comments(bid, aid):
    """
    swagger_from_file: apidoc/load_article_comments.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/comments/first'))
def _load_article_firstcomments(bid, aid):
    """
    swagger_from_file: apidoc/load_article_firstcomments.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>'))
def _get_board_detail(bid):
    """
    swagger_from_file: apidoc/get_board_detail.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/summary'))
def _get_board_summary(bid):
    """
    swagger_from_file: apidoc/get_board_summary.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/users'))
def _get_board_users(bid):
    """
    swagger_from_file: apidoc/load_board_users.yaml

    """
    return ''
