import os
from jinja2 import Environment, FileSystemLoader, StrictUndefined, select_autoescape
from utils.migration import read_conf

revision = '1.10.0'
down_revisions = ['1.9.0']

def migrate(input_cfg, output_cfg):
    config_dict = read_conf(input_cfg)

    current_dir = os.path.dirname(__file__)
    tpl = Environment(
        loader=FileSystemLoader(current_dir),
        undefined=StrictUndefined,
        trim_blocks=True,
        lstrip_blocks=True,
        autoescape = select_autoescape()
        ).get_template('harbor.yml.jinja')

    with open(output_cfg, 'w') as f:
        f.write(tpl.render(**config_dict))