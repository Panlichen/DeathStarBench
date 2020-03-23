after install `jinja2` via  `pip install jinja2-cli`

``./generate_secrets.sh all `./generate_secrets.sh fsid` ``fails because:
```
Traceback (most recent call last):
  File "/usr/local/bin/jinja2", line 8, in <module>
    sys.exit(main())
  File "/usr/local/lib/python3.5/dist-packages/jinja2cli/cli.py", line 424, in main
    sys.exit(cli(opts, args))
  File "/usr/local/lib/python3.5/dist-packages/jinja2cli/cli.py", line 314, in cli
    out.write(render(template_path, data, extensions, opts.strict))
  File "/usr/local/lib/python3.5/dist-packages/jinja2cli/cli.py", line 229, in render
    return env.get_template(os.path.basename(template_path)).render(data)
  File "/usr/local/lib/python3.5/dist-packages/jinja2/environment.py", line 1008, in render
    return self.environment.handle_exception(exc_info, True)
  File "/usr/local/lib/python3.5/dist-packages/jinja2/environment.py", line 780, in handle_exception
    reraise(exc_type, exc_value, tb)
  File "/usr/local/lib/python3.5/dist-packages/jinja2/_compat.py", line 37, in reraise
    raise value.with_traceback(tb)
  File "/home1/root/DeathStarBench/ceph-deploy/generator/templates/ceph/ceph.conf.jinja", line 3, in top-level template code
    cephx = {{ environ("auth_cephx", "true") }}
TypeError: <lambda>() takes 1 positional argument but 2 were given
```