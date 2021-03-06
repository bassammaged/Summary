module.exports = {
  define: {
    __EXP__: '1 + 1',
    __STRING__: '"hello"',
    __NUMBER__: 123,
    __BOOLEAN__: true,
    __OBJ__: {
      foo: 1,
      bar: {
        baz: 2
      },
      process: {
        env: {
          SOMEVAR: '"PROCESS MAY BE PROPERTY"'
        }
      }
    },
    __VAR_NAME__: false,
    'process.env.SOMEVAR': '"SOMEVAR"',
    __IMPORT_FILE_NAME__: '"importFileName"',
    __EXPORT_FILE_NAME__: '"exportFileName"',
    PATH: '"filePath"'
  }
}
