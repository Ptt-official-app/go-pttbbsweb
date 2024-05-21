#!/usr/bin/env python
# -*- coding: utf-8 -*-

import yaml
import sys
import re
import json

filename = sys.argv[1]

with open(filename, 'r') as f:
    the_struct = yaml.full_load(f)

components = the_struct.get('components', {})
components['securitySchemes'] = {
    'default': {
        'type': 'http',
        'scheme': 'bearer',
    },
}

to_update = {
    'components': components,
    'security': [
        {
            'default': [],
        },
    ]
}

the_struct.update(to_update)

out_filename = re.sub(r'\.yaml$', '.security.json', filename)
with open(out_filename, 'w') as f:
    json.dump(the_struct, f, indent=4)
