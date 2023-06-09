#include <iostream>
#include <google/protobuf/compiler/plugin.h>
#include <google/protobuf/compiler/code_generator.h>
#include <google/protobuf/io/zero_copy_stream.h>
#include <google/protobuf/io/zero_copy_stream_impl.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/descriptor.pb.h>
#include <google/protobuf/descriptor.h>

class SolidityGenerator : public google::protobuf::compiler::CodeGenerator
{
public:
    bool Generate(const google::protobuf::FileDescriptor *file, const std::string &parameter,
                  google::protobuf::compiler::GeneratorContext *context, std::string *error) const
    {
        // show all messages got from protoc
        for (int i = 0; i < file->message_type_count(); ++i)
        {
            const google::protobuf::Descriptor *message = file->message_type(i);
            std::cout << message->name() << std::endl;

            // show fields in message
            for (int j = 0; j < message->field_count(); ++j)
            {
                const google::protobuf::FieldDescriptor *field = message->field(j);
                std::cout << "field type: " << field->type_name() << std::endl;
                std::cout << "field name: " << field->name() << std::endl;
                std::cout << "field nuber: " << field->number() << std::endl;
            }
        }

        return true;
    }
};

int main(int argc, char *argv[])
{
    SolidityGenerator generator;
    return google::protobuf::compiler::PluginMain(argc, argv, &generator);
}
