import gspread
import pandas as pd
import matplotlib.pyplot as plt
from oauth2client.service_account import ServiceAccountCredentials
import numpy as np
from fpdf import FPDF
from matplotlib.backends.backend_pdf import PdfPages
import statistics



"""
Esta funcion lee la hoja de calculo en tiempo real
"""
def leer_hoja_calculo():
    # Cargar las credenciales del archivo JSON
    scope = ['https://www.googleapis.com/auth/spreadsheets']
    creds = ServiceAccountCredentials.from_json_keyfile_name('key.json', scope)

    # Autenticar la conexión
    client = gspread.authorize(creds)

    # Abrir la hoja de Google Sheets por su clave única
    sheet = client.open_by_key('1baeM9y-PVOgMQLxmVwiQBcYA6yl05yPRAFwErRlDNOg')

    # Leer los datos de la primera hoja de la hoja de Google Sheets
    worksheet = sheet.sheet1
    data = worksheet.get_all_values()
    df = pd.DataFrame(data)

    # ruta_archivo = '/home/jorge/Documentos/graficos_analisis/archivo.xlsx'

    # Imprimir los datos

    # df.drop(df.index[0:9], inplace=True)
    df=df.drop(df.index[0:8])
    df = df.drop(df.columns[0:3], axis='columns')

    return df
"""
Esta funcion obtiene el promeido de timepo de ejecucion de cada algoritmo
Lo devuelve en un arreglo bidimensional
"""
def obtenerPromedioEjecucionAlgoritmos(df):

    

    valores_metodos=[]

    

    for j in range (len(df.index)):



        f1 = df.iloc[j]
        
        f1=f1.drop(f1.index[0])

        valores_algoritmo = 0.0000000

        for i, elemento in enumerate(f1):
                
            if not(i % 2 == 0):       
               
                try:
                    #Esta linea de codigo cambia el formato de cada numero , 
                    #esta documentada porque hay un metood que ya lo hace antes que este
                    # elemento = elemento.replace(",", ".")
                    valores_algoritmo = valores_algoritmo+ float(elemento)

                except ValueError as e:
                    print(f"No se pudo convertir el valor a entero: {elemento}")
            

        valores_metodos.append(valores_algoritmo)
    return valores_metodos   

"""
Esta funcion genera un grafico de barras con los tiempos de ejecucion de cada algoritmo
"""
def generar_grafico_barras_promedio_tiempo_algoritmos(tiempo_ejecucion):

    etiquetas_metodos = ["NaivStandard", "NaivOnArray", "NaivKahan", "NaivLoopUnrollingTwo", "NaivLoopUnrollingThree", "NaivLoopUnrollingFour", "WinogradOriginal", "WinogradScaled", "StrassenNaiv", "StrassenWinograd", "III.3 Sequential block", "III.4 Parallel Block", "IV.3 Sequential block", "IV.4 Parallel Block", "V.3 Sequential block", "V.4 Parallel Block"]





    plt.barh(etiquetas_metodos, tiempo_ejecucion)
    plt.xlabel('Tiempo en segundos')
    plt.ylabel('Algoritmos')
    plt.title('Tiempo de ejecucion por algoritmo')
    plt.gcf().set_size_inches(10.4, 8.4)
    plt.savefig('graficos/grafico_tiempos_promedio_algoritmos.png')
    plt.clf()  

"""
Esta funcion genera un grafico de barras con los tiempos de ejecucion de cada algoritmo
Ordenando de manera descendente
"""
def generar_grafico_barras_promedio_tiempo_algoritmos_ascendente(tiempo_ejecucion):

    etiquetas_metodos = ["NaivStandard", "NaivOnArray", "NaivKahan", "NaivLoopUnrollingTwo", "NaivLoopUnrollingThree", "NaivLoopUnrollingFour", "WinogradOriginal", "WinogradScaled", "StrassenNaiv", "StrassenWinograd", "III.3 Sequential block", "III.4 Parallel Block", "IV.3 Sequential block", "IV.4 Parallel Block", "V.3 Sequential block", "V.4 Parallel Block"]

 



    diccionario = dict(zip(etiquetas_metodos,tiempo_ejecucion))
    diccionario_ordenado = dict(sorted(diccionario.items(), key=lambda item: item[1]))

    claves = list(diccionario_ordenado.keys())
    valores = list(diccionario_ordenado.values())



    plt.barh(claves, valores)
 
    plt.xlabel('Tiempo en segundos')
    plt.ylabel('Algoritmos')
    plt.title('Tiempo de ejecucion por algoritmo')
    plt.savefig('graficos/grafico_tiempos_promedio_algoritmos_ordenados.png')


"""
Esta funcion obtiene los tiempo del dataframe , para una mejor manipulacion
y Realiza la conversion a float 
"""
def obtener_datos_en_tabla(df):



    valores_metodos=[]


    #Itera sobre todos los algoritmos
    datos_tiempo_algoritmo =  []

    datos_tiempo_todos_los_algoritmos =  []

    for j in range (len(df.index)):


        #Obtiene la primera fila del df
        f1 = df.iloc[j]
        #Le quita el nombre del algoritmo
        f1=f1.drop(f1.index[0])

        valor_en_decimal = 0.0000000
        #Itero sobre los datos de cada algoritmo
        for i, elemento in enumerate(f1):
                
            if (i % 2 == 0):      

               
                try:
                    elemento = elemento.replace(",", ".")
                    valor_en_decimal= float(elemento)
                    datos_tiempo_algoritmo.append(valor_en_decimal)

                except ValueError as e:
                    print(f"No se pudo convertir el valor a entero: {elemento}")
    datos_tiempo_todos_los_algoritmos.append(datos_tiempo_algoritmo)

    datos_tiempo_todos_los_algoritmos=np.array(datos_tiempo_todos_los_algoritmos).reshape(16,12)



    return datos_tiempo_todos_los_algoritmos
"""
Esta funcion genera un pdf , dado un dataframe
"""
def generar_pdf(df):
    with PdfPages('graficos/medidas_estadistica_descriptiva.pdf') as pdf:
        fig = plt.figure(figsize=(11.69,8.27), dpi=100)
        ax = fig.add_subplot(111)
        ax.table(cellText=df.values,colLabels=df.columns,loc='center')
        ax.axis('off')
        pdf.savefig(fig)

"""
Esta funcion itera sobre cada algoritmo y obtiene las medidas de tendencia central
"""
def obtener_medidas_tendencia_central(datos_algoritmos):

    medidas_tendecia_central=[]
    

    for fila in datos_algoritmos:

            
            media = statistics.mean(fila) 
            desviacion_estandar = statistics.stdev(fila) 
            varianza = statistics.variance(fila) 
            rango = max(fila) - min(fila) # 12
            medidas_tendecia_central.append(media)
            medidas_tendecia_central.append(desviacion_estandar)
            medidas_tendecia_central.append(varianza)
            medidas_tendecia_central.append(rango)
    medidas_tendecia_central=np.array(medidas_tendecia_central).reshape(16,4)
   
    return medidas_tendecia_central

"""
Esta funcion le pone las etiquetas al datafram que se convertira en la tabla en un pdf
"""
def armar_tabla(df):

    nombres_medidas_tendencia_central=["Nombre","Media","Rango","Desviacion estandar","Varianza"]
    etiquetas_metodos = ["NaivStandard", "NaivOnArray", "NaivKahan", "NaivLoopUnrollingTwo", "NaivLoopUnrollingThree", "NaivLoopUnrollingFour", "WinogradOriginal", "WinogradScaled", "StrassenNaiv", "StrassenWinograd", "III.3 Sequential block", "III.4 Parallel Block", "IV.3 Sequential block", "IV.4 Parallel Block", "V.3 Sequential block", "V.4 Parallel Block"]
    df.insert(0, '0', etiquetas_metodos)
    df.loc[-1] = nombres_medidas_tendencia_central
    df.index = df.index + 1
    df = df.sort_index()

    return df
 
    


"""
Esta funcion graficas los ttiempos de ejecucion de cada uno de los algoritmos respecto a su tamaño
"""
def generar_graficas(df):
    etiquetas_metodos = ["NaivStandard", "NaivOnArray", "NaivKahan", "NaivLoopUnrollingTwo", "NaivLoopUnrollingThree", "NaivLoopUnrollingFour", "WinogradOriginal", "WinogradScaled", "StrassenNaiv", "StrassenWinograd", "III.3 Sequential block", "III.4 Parallel Block", "IV.3 Sequential block", "IV.4 Parallel Block", "V.3 Sequential block", "V.4 Parallel Block"]




    for j in range (len(df.index)):


        valores_cuadrados_matriz=["2","4","8","16","32","64","128","256", "512", "1024", "2048", "4096"]




        valores_tiempo=df.iloc[j]


        plt.title(''+etiquetas_metodos[j])
        plt.xlabel('Tamaño de las matrices')
        plt.ylabel('Tiempo de ejecución (s)')



        # Crear gráfico
        plt.plot(valores_cuadrados_matriz, valores_tiempo,'o-')


        
        plt.savefig('graficos/grafico_'+etiquetas_metodos[j]+'.png')
        plt.clf()  
        # plt.show()



if __name__ == "__main__":
    df=leer_hoja_calculo()


    datos_todos_los_algoritmos=obtener_datos_en_tabla(df)




    dfn = pd.DataFrame(datos_todos_los_algoritmos)

    tiempo_ejecucion=obtenerPromedioEjecucionAlgoritmos(dfn)
    tiempo_2 = tiempo_ejecucion.copy()
    generar_grafico_barras_promedio_tiempo_algoritmos(tiempo_ejecucion)
    
    generar_grafico_barras_promedio_tiempo_algoritmos_ascendente(tiempo_2)


    #Se llama a la funcion que genera las graficas 
    generar_graficas(dfn)



    #Se llama a la funcion que obtiene las medias de tendecia central
    dfn=obtener_medidas_tendencia_central(datos_todos_los_algoritmos)
    dfp = pd.DataFrame(dfn)
    
    dfp=armar_tabla(dfp)
    # #Se genera un pdf con las medidas
    generar_pdf(dfp)





 


            


           


    