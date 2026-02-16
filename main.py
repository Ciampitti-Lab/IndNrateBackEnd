from flask import Flask, Blueprint, render_template, request, jsonify
import geopandas as gpd
import pandas as pd
import plotly.graph_objects as go
import numpy as np
import psycopg2
import os

main = Blueprint('main', __name__)

@main.route("/")
def main_interface():
    return render_template('indianaaonr.html')


@main.route("/generate_fig")
def generate_fig():
    cell = int(request.args.get('cell'))
    fig_creation_aonr(cell)
    return jsonify({"url": "/static/images/fig.html"})

@main.route("/generate_eonr_fig")
def generate_eonr_fig():
    cell = int(request.args.get('cell'))
    grainPrice = int(request.args.get('grain_price'))
    nPrice = float(request.args.get('n_price'))
    fig_creation_eonr(cell,grainPrice,nPrice)
    return jsonify({"url": "/static/images/fig2.html"})

# DB information (Render Database)
DB_CONFIG = {
    'host': 'dpg-d69j6cpr0fns7384trhg-a.oregon-postgres.render.com', 
    'dbname': 'ind_rates_apsim_database',
    'user': 'ind_rates_apsim_database_user',
    'password': 'QmOUyd5ZIWbVctqd79rQvNCFLTEAZBst',
    'port': 5432
}

def fig_creation_aonr(cell):
    conn = psycopg2.connect(**DB_CONFIG)

    simulations = pd.read_sql("SELECT * FROM simulations;", conn)

    conn.close()

    simulations['yield_kg_ha'] = simulations['yield_kg_ha']

    sim_cell=simulations[simulations['id_cell']==cell]
    max_row = sim_cell.loc[sim_cell ['yield_kg_ha'].idxmax()]


    x_smooth = sim_cell['nitro_kg_ha']
    y_smooth = sim_cell['yield_kg_ha']


    fig = go.Figure()


    fig.add_trace(go.Scatter(
        x=x_smooth,
        y=y_smooth,
        mode='lines',
        name='Yield Curve',
        line=dict(color='steelblue', width=3),
    ))


    fig.add_trace(go.Scatter(
        x=[max_row['nitro_kg_ha']],
        y=[max_row['yield_kg_ha']],
        mode='markers+text',
        name='Max Yield',
        marker=dict(color='red', size=12)
    ))


    fig.add_shape(
        type='line',
        x0=max_row['nitro_kg_ha'],
        y0=max_row['yield_kg_ha'],
        x1=max_row['nitro_kg_ha'],
        y1=max_row['yield_kg_ha'],
        line=dict(color='green', width=3, dash='dash'),
        name='Max Line'
    )

    fig.update_layout(
        xaxis_title='Nitrogen (kg/ha)',
        yaxis_title='Yield (t/ha)',
        plot_bgcolor='white',
        xaxis=dict(
            showgrid=True,
            gridcolor='lightgrey',
            zeroline=False,
            showline=True,
            linecolor='black',
        ),
        yaxis=dict(
            showgrid=True,
            gridcolor='lightgrey',
            zeroline=False,
            showline=True,
            linecolor='black',
        ),
        legend=dict(
            bgcolor='rgba(0,0,0,0)',
            bordercolor='rgba(0,0,0,0)',
        )
    )
    

    fig.write_html('static/images/fig.html')


def fig_creation_eonr(cell,grainPrice,nPrice):
    conn = psycopg2.connect(**DB_CONFIG)

    simulations = pd.read_sql("SELECT * FROM simulations;", conn)

    conn.close()

    df_cell=simulations[simulations['id_cell']==cell]


    df_cell = simulations[simulations['id_cell'] == cell]


    df_cell['economic'] = df_cell['yield_kg_ha'] * grainPrice - df_cell['nitro_kg_ha'] * nPrice


    max_row = df_cell.loc[df_cell['economic'].idxmax()]

    x_smooth = df_cell['nitro_kg_ha']
    y_smooth = df_cell['economic']


    fig = go.Figure()


    fig.add_trace(go.Scatter(
        x=x_smooth,
        y=y_smooth,
        mode='lines',
        name='Economic Return Curve',
        line=dict(color='steelblue', width=3),
    ))


    fig.add_trace(go.Scatter(
        x=[max_row['nitro_kg_ha']],
        y=[max_row['economic']],
        mode='markers+text',
        name='EONR',
        marker=dict(color='red', size=12)
    ))


    fig.add_shape(
        type='line',
        x0=max_row['nitro_kg_ha'],
        y0=max_row['economic'],
        x1=max_row['nitro_kg_ha'],
        y1=max_row['economic'],
        line=dict(color='green', width=3, dash='dash')
    )


    fig.update_layout(
        xaxis_title='Nitrogen (kg/ha)',
        yaxis_title='Economic Return (USD/ha)',
        plot_bgcolor='white',
        xaxis=dict(showgrid=True, gridcolor='lightgrey', zeroline=False, showline=True, linecolor='black'),
        yaxis=dict(showgrid=True, gridcolor='lightgrey', zeroline=False, showline=True, linecolor='black'),
        legend=dict(bgcolor='rgba(0,0,0,0)', bordercolor='rgba(0,0,0,0)')
    )
    
    fig.write_html('static/images/fig2.html')
