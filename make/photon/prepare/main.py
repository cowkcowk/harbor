from commands.prepare import prepare

import click

@click.group()
def cli():
    pass

cli.add_command(prepare)

if __name__ == '__main__':
    cli()
