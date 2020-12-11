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


@app.route(_with_app_prefix('/Account/login'), methods=['POST'])
def _login():
    """
    swagger_from_file: apidoc/account_login.yaml

    """
    return ''


@app.route(_with_app_prefix('/Account/register'), methods=['POST'])
def _register():
    """
    swagger_from_file: apidoc/account_register.yaml

    """
    return ''


@app.route(_with_app_prefix('/Article/populars'))
def _get_popular_post():
    """
    swagger_from_file: apidoc/get_popular_post.yaml

    """
    return ''


@app.route(_with_app_prefix('/Board/favorite/<username>'))
def _get_favorite_board(username):
    """
    swagger_from_file: apidoc/get_favorite_board.yaml

    """
    return ''


@app.route(_with_app_prefix('/Board/search'))
def _find_board_by_name():
    """
    swagger_from_file: apidoc/find_board_by_name.yaml
    """
    return ''


@app.route(_with_app_prefix('/Board/populars'))
def _get_popular_board_list():
    """
    swagger_from_file: apidoc/get_popular_board_list.yaml

    """
    return ''


@app.route(_with_app_prefix('/User/Users/<username>'))
def _get_user_info(username):
    """
    swagger_from_file: apidoc/get_user_info.yaml

    """
    return ''


@app.route(_with_app_prefix('/User/Article/<username>'))
def _get_user_post_list(username):
    """
    swagger_from_file: apidoc/get_user_post_list.yaml

    Args:
        username (TYPE): Description

    Returns:
        TYPE: Description
    """
    return ''


@app.route(_with_app_prefix('/User/Comment/<username>'))
def _get_user_comment_list(username):
    """
    swagger_from_file: apidoc/get_user_comment_list.yaml

    """
    return ''


@app.route(_with_app_prefix('/Article/Articles/<board>'))
def _get_post_list(board):
    """
    swagger_from_file: apidoc/get_post_list.yaml

    """
    return ''


@app.route(_with_app_prefix('/Article/Articles/<board>/<article>'))
def _get_post(board, article):
    """
    swagger_from_file: apidoc/get_post.yaml

    """
    return ''


@app.route(_with_app_prefix('/Board/Boards/<board>'))
def _get_board_detail(board):
    """
    swagger_from_file: apidoc/get_board_detail.yaml

    """
    return ''


@app.route(_with_app_prefix('/Board/summary/<board>'))
def _get_board_title(board):
    """
    swagger_from_file: apidoc/get_board_summary.yaml

    """
    return ''
