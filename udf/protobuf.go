package udf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/msgregistry"
	"github.com/pkg/errors"
	"github.com/viant/toolbox"
	"github.com/viant/toolbox/data"
	"io"
	"io/ioutil"
	"path"
	"strings"
)

//ProtoCodec represent a proto codec
type ProtoCodec struct {
	registry            *msgregistry.MessageRegistry
	msgType             string
	convertToLowerCamel bool
}

func (c *ProtoCodec) AsMessage(msgType string, data []byte) (interface{}, error) {
	msgDescriptor, err := c.registry.FindMessageTypeByUrl(msgType)
	if err != nil {
		return nil, err
	}
	protoMsg := dynamic.NewMessage(msgDescriptor)
	if err = protoMsg.Unmarshal(data); err != nil {
		return nil, err
	}
	JSON, err := protoMsg.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var result = make(map[string]interface{})
	err = toolbox.NewJSONDecoderFactory().Create(bytes.NewReader(JSON)).Decode(&result)
	return result, err
}

func (c *ProtoCodec) AsBinary(msgType string, msg interface{}) ([]byte, error) {
	msgDescriptor, err := c.registry.FindMessageTypeByUrl(msgType)
	if err != nil {
		return nil, err
	}
	var reader io.Reader
	switch value := msg.(type) {
	case string:
		reader = strings.NewReader(value)
	case []byte:
		reader = bytes.NewReader(value)
	default:
		text, err := toolbox.AsJSONText(msg)
		if err != nil {
			return nil, err
		}
		reader = strings.NewReader(text)
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if c.convertToLowerCamel {
		data, err = c.toLowerCamel(err, data)
		if err != nil {
			err = errors.Wrapf(err, "failed to convert to lowerCase fields")
			return nil, err
		}
	}
	protoMsg := dynamic.NewMessage(msgDescriptor)
	err = protoMsg.UnmarshalJSON(data)
	if err != nil {
		err = errors.Wrapf(err, "failed to UnmarshalJSON")
		return nil, err
	}
	return protoMsg.Marshal()
}


func (c *ProtoCodec) toLowerCamel(err error, data []byte) ([]byte, error) {
	aMap := map[string]interface{}{}
	err = json.Unmarshal(data, &aMap)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid JSON")
	}
	transformed := map[string]interface{}{}
	err = toolbox.CopyMap(aMap, transformed, func(key, value interface{}) (interface{}, interface{}, bool) {
		if value == nil || toolbox.AsString(value) == "" || key == nil {
			return nil, nil, false
		}
		return toolbox.ToCaseFormat(toolbox.AsString(key), toolbox.CaseUpperCamel, toolbox.CaseLowerCamel), value, true
	})
	if err != nil {
		return nil, err
	}
	return json.Marshal(transformed)
}


//NewProtoCodec creates a new protobuf codec
func NewProtoCodec(schemaFile, importPath string, msgType string, lowercaseKey bool) (*ProtoCodec, error) {
	parser := protoparse.Parser{ImportPaths: []string{importPath}, IncludeSourceCodeInfo: true}
	descriptors, err := parser.ParseFiles(schemaFile)
	if err != nil {
		return nil, err
	}
	baseURL := ""
	registry := msgregistry.NewMessageRegistryWithDefaults()
	for _, desc := range descriptors {
		registry.AddFile(baseURL, desc)
	}
	return &ProtoCodec{
		registry:            registry,
		msgType:             msgType,
		convertToLowerCamel: lowercaseKey,
	}, nil

}

func getProtoCodec(source string, args []interface{}) (*ProtoCodec, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("no sufficent args |usage: %v(schemaFile, messageType, importPath)", source)
	}
	schemaFile := toolbox.AsString(args[0])
	messageType := toolbox.AsString(args[1])
	importPath, filename := path.Split(schemaFile)
	if len(args) > 2 && args[2] != nil  {
		importPath = toolbox.AsString(args[2])
	} else {
		schemaFile = filename
	}
	lowercaseKey := false
	if len(args) > 3 {
		lowercaseKey = toolbox.AsBoolean(args[3])
	}
	return NewProtoCodec(schemaFile, importPath, messageType, lowercaseKey)
}

//NewProtoWriter creates a new proto writer provider
func NewProtoWriter(args ...interface{}) (func(source interface{}, state data.Map) (interface{}, error), error) {
	codec, err := getProtoCodec("NewProtoWriter", args)
	if err != nil {
		return nil, err
	}
	return func(source interface{}, state data.Map) (interface{}, error) {
		data, err := codec.AsBinary(codec.msgType, source)
		return data, err
	}, nil
}

//NewProtoWriter creates a new proto writer provider
func NewProtoReader(args ...interface{}) (func(source interface{}, state data.Map) (interface{}, error), error) {
	codec, err := getProtoCodec("NewProtoReader", args)
	if err != nil {
		return nil, err
	}
	return func(source interface{}, state data.Map) (interface{}, error) {
		var reader io.Reader
		switch val := source.(type) {
		case []byte:
			reader = bytes.NewReader(val)
		case string:
			reader = strings.NewReader(val)
		case io.Reader:
			reader = val
		}
		binaryData, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		return codec.AsMessage(codec.msgType, binaryData)
	}, nil
}
